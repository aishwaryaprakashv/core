package migration

import (
	"database/sql"
	"log"

	v1 "github.com/onepanelio/core/pkg"
	uid2 "github.com/onepanelio/core/pkg/util/uid"
	"github.com/pressly/goose"
)

const tensorflowWorkflowTemplate = `entrypoint: main
arguments:
    parameters:
    - name: source
      value: https://github.com/onepanelio/tensorflow-examples.git
    - name: command
      value: "python mnist/main.py --epochs=5"
volumeClaimTemplates:
  - metadata:
      name: data
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 2Gi
  - metadata:
      name: output
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 2Gi
templates:
  - name: main
    dag:
      tasks:
      - name: train-model
        template: pytorch
# Uncomment section below to send metrics to Slack
#      - name: notify-in-slack
#        dependencies: [train-model]
#        template: slack-notify-success
#        arguments:
#          parameters:
#          - name: status
#            value: "{{tasks.train-model.status}}"
#          artifacts:
#          - name: metrics
#            from: "{{tasks.train-model.outputs.artifacts.sys-metrics}}"
  - name: pytorch
    inputs:
      artifacts:
      - name: src
        path: /mnt/src
        git:
          repo: "{{workflow.parameters.source}}"
    outputs:
      artifacts:
      - name: model
        path: /mnt/output
        optional: true
        archive:
          none: {}
    container:
      image: tensorflow/tensorflow:latest
      command: [sh,-c]
      args: ["{{workflow.parameters.command}}"]
      workingDir: /mnt/src
      volumeMounts:
      - name: data
        mountPath: /mnt/data
      - name: output
        mountPath: /mnt/output
  - name: slack-notify-success
    container:
      image: technosophos/slack-notify
      command: [sh,-c]
      args: ['SLACK_USERNAME=Worker SLACK_TITLE="{{workflow.name}} {{inputs.parameters.status}}" SLACK_ICON=https://www.gravatar.com/avatar/5c4478592fe00878f62f0027be59c1bd SLACK_MESSAGE=$(cat /tmp/metrics.json)} ./slack-notify']
    inputs:
      parameters:
      - name: status
      artifacts:
      - name: metrics
        path: /tmp/metrics.json
        optional: true
`

const tensorflowWorkflowTemplateName = "TensorFlow Training"

func initialize20200605090535() {
	if _, ok := initializedMigrations[20200605090535]; !ok {
		goose.AddMigration(Up20200605090535, Down20200605090535)
		initializedMigrations[20200605090535] = true
	}
}

// Up20200605090535 will insert a tensorflow workflow template to each user.
// Each user is determined by onepanel enabled namespaces.
// Any errors reported are logged as fatal.
func Up20200605090535(tx *sql.Tx) error {
	client, err := getClient()
	if err != nil {
		return err
	}
	defer client.DB.Close()

	migrationsRan, err := getRanSQLMigrations(client)
	if err != nil {
		return err
	}

	if _, ok := migrationsRan[20200605090535]; ok {
		return nil
	}

	namespaces, err := client.ListOnepanelEnabledNamespaces()
	if err != nil {
		return err
	}

	workflowTemplate := &v1.WorkflowTemplate{
		Name:     tensorflowWorkflowTemplateName,
		Manifest: tensorflowWorkflowTemplate,
	}

	for _, namespace := range namespaces {
		if _, err := client.CreateWorkflowTemplate(namespace.Name, workflowTemplate); err != nil {
			return err
		}
	}
	return nil
}

// Down20200605090535 will attempt to remove tensorflow workflow from each user.
// Each user is determined by onepanel enabled namespaces.
// DB entries are archived, K8S components are deleted.
// Active workflows with that template are terminated.
// Any errors reported are logged as fatal.
func Down20200605090535(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	client, err := getClient()
	if err != nil {
		return err
	}

	namespaces, err := client.ListOnepanelEnabledNamespaces()
	if err != nil {
		return err
	}

	uid, err := uid2.GenerateUID(tensorflowWorkflowTemplateName, 30)
	if err != nil {
		return err
	}
	for _, namespace := range namespaces {
		if _, err := client.ArchiveWorkflowTemplate(namespace.Name, uid); err != nil {
			log.Fatalf("error %v", err.Error())
		}
	}

	return nil
}
