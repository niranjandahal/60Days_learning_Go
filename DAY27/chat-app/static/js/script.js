// static/js/script.js

document.addEventListener("DOMContentLoaded", () => {
    fetchMessages();
});

function fetchMessages() {
    fetch('/messages')
        .then(response => response.json())
        .then(messages => {
            const chatWindow = document.getElementById('chatwindow');
            chatWindow.innerHTML = '';
            messages.forEach(message => {
                const messageElement = document.createElement('div');
                messageElement.textContent = `${message.created_at} - ${message.username}: ${message.content}`;
                chatWindow.appendChild(messageElement);
            });
        })
        .catch(error => console.error('Error fetching messages:', error));
}

function sendMessage() {
    const messageInput = document.getElementById('message');
    const content = messageInput.value;
    const username = "your_username";  // Replace with the actual username logic

    if (content) {
        fetch('/message', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ content, username })
        })
        .then(response => {
            if (response.ok) {
                messageInput.value = '';
                fetchMessages();
            } else {
                console.error('Failed to send message');
            }
        })
        .catch(error => console.error('Error sending message:', error));
    }
}
