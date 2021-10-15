import React, { Component } from "react";
import Message from "../message/Message"
import "./Chat.scss";

class Chat extends Component {
    render() {
        console.log(this.props.chat);

        const messages = this.props.chat.map(msg => <Message message={msg.data} />);

        return (
            <div className='Chat'>
                <h2>Chat History</h2>
                {messages}
            </div>
        );
    };
}

export default Chat;