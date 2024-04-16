echo "#### init example_01.sh ####"


# Defina o endpoint para se comunicar com o localstack
ENDPOINT_URL=http://localstack:4566

# Cria um tópico SNS chamado "command"
TOPIC_ARN=$(aws --endpoint-url=${ENDPOINT_URL} sns create-topic --name command --query 'TopicArn' --output text)

echo "Topic ARN: ${TOPIC_ARN}"

# Cria uma fila SQS chamada "create_example"
QUEUE_URL=$(aws --endpoint-url=${ENDPOINT_URL} sqs create-queue --queue-name create_example --query 'QueueUrl' --output text)

echo "Queue URL: ${QUEUE_URL}"

# Obtém o ARN da fila SQS
QUEUE_ARN=$(aws --endpoint-url=${ENDPOINT_URL} sqs get-queue-attributes --queue-url ${QUEUE_URL} --attribute-names QueueArn --query 'Attributes.QueueArn' --output text)

echo "Queue ARN: ${QUEUE_ARN}"

# Assina o tópico SNS para a fila SQS
SUBSCRIPTION_ARN=$(aws --endpoint-url=${ENDPOINT_URL} sns subscribe --topic-arn ${TOPIC_ARN} --protocol sqs --notification-endpoint ${QUEUE_ARN} --query 'SubscriptionArn' --output text)

echo "Subscription ARN: ${SUBSCRIPTION_ARN}"

# Define o atributo de filtro para a assinatura
aws --endpoint-url=${ENDPOINT_URL} sns set-subscription-attributes --subscription-arn ${SUBSCRIPTION_ARN} --attribute-name FilterPolicy --attribute-value '{"Type":["CREATE"]}'
echo "#### end example_01.sh ####"