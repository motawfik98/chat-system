class MessagesController < ActionController::API
  def index
    app_token = params[:token]
    chat_number = params[:number]
    chat = Chat.where(app_id: Application.where(token: app_token), number: chat_number).first
    if chat
      message = params[:message]
      operator = params[:operator]
      if operator != "and" && operator != "or"
        render json: { status: 400, message: "invalid option for operator field. valid: (and/or)" }, status: :bad_request
      else
        response = Message.search(message, chat.id, operator)

        results = response.results
        messages = results.map do |message|
          message._source.appToken = app_token
          message._source.delete(:chat_id)
          message._source # return actual message object
        end
        render json: messages.as_json(methods: [:appToken])
      end
    else
      render json: { status: 404, message: "cannot find specified app and chat combination" }, status: :not_found
    end
  end
end
