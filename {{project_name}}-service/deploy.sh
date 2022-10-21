REGION=europe-west2
GATEWAY_URL=https://android-gateway-clzdlli7.nw.gateway.dev

gcloud run deploy {{project_name}}-service \
  --source . \
  --region=$REGION \
  --use-http2 \
  --set-env-vars GATEWAY_URL=$GATEWAY_URL