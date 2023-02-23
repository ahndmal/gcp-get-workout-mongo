deploy:
	gcloud functions deploy get-workout-go \
		--gen2 \
		--region=us-central1 \
		--trigger-http \
		--memory=128Mi \
		--source=. \
        --runtime=go120 \
        --entry-point=GetWorkout \
        --allow-unauthenticated