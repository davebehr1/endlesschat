package com.chatserver;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.messaging.handler.annotation.SendTo;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ChatController {

    @Autowired
    SimpMessagingTemplate template;

    @PostMapping("/send")
    public ResponseEntity<Void> sendMessage(@RequestBody Message message) {
        template.convertAndSend("/topic/greetings", message);
        return new ResponseEntity<>(HttpStatus.ACCEPTED.OK);
    }

    @MessageMapping("/chat")
    public ResponseEntity<Void> receiveMessage(@Payload Message message) {
        System.out.println(message.getName());
        template.convertAndSend("/topic/greetings", message);
        return new ResponseEntity<>(HttpStatus.ACCEPTED.OK);

    }
    @SendTo("/topic/greetings")
    public Message broadcastMessage(@Payload  Message message) {
        return message;
    }

}
