deploy:
	~/google-cloud-sdk/bin/gcloud functions deploy GetWorkout --trigger-http --runtime=go116 --entry-point=GetWorkout --allow-unauthenticated --memory=256MB