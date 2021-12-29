require 'elasticsearch/model'

class Message < ApplicationRecord
  include Elasticsearch::Model
  include Elasticsearch::Model::Callbacks

  def self.search(message, chat_id, operator)
    __elasticsearch__.search(
      {
        query: {
          bool: {
            must: [
              {
                match: {
                  message: {
                    query: message,
                    operator: operator
                  }
                }
              },
              {
                match: {
                  chat_id: chat_id
                }
              }
            ]
          }
        }
      }
    )
  end
end

