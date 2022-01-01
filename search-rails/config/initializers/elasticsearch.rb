config = {
  elasticsearch_host: ENV["ELASTICSEARCH_URL"]
}

Elasticsearch::Model.client = Elasticsearch::Client.new(config)
