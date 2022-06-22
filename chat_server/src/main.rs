use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};

#[get("/")]
async fn index() -> impl Responder {
    HttpResponse::Ok().body("Rust Chat Server")
}

#[get("/healthcheck")]
async fn healthcheck() -> impl Responder {
    HttpResponse::Ok().body("I'm alive!")
}

pub fn init(config: &mut web::ServiceConfig) {
    config.service(web::scope("").service(index).service(healthcheck));
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("Starting web server");
    HttpServer::new(|| App::new().configure(init))
        .bind("0.0.0.0:5003")?
        .run()
        .await
}
