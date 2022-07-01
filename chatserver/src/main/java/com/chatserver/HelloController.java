package com.chatserver;


import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.stereotype.Controller;
import org.springframework.messaging.handler.annotation.SendTo;
import org.springframework.messaging.handler.annotation.MessageMapping;

@RestController
public class HelloController {
    @GetMapping("/")
    public String index() {
       return "Chat Server";
    }
}

