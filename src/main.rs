#![deny(clippy::all)]

use actix_web::{
    middleware::{Compress, NormalizePath},
    App, HttpServer,
};
use miette::{IntoDiagnostic, Result};
use tracing_actix_web::TracingLogger;

#[tokio::main]
async fn main() -> Result<()> {
    dotenv::dotenv().ok();
    tracing_subscriber::fmt::init();
    let host = std::env::var("HOST").unwrap_or_else(|_| "127.0.0.1".to_string());
    let port = std::env::var("PORT").unwrap_or_else(|_| "7000".to_string());
    let addr = format!("{}:{}", host, port);
    HttpServer::new(|| {
        App::new()
            .wrap(TracingLogger::default())
            .wrap(Compress::default())
            .wrap(NormalizePath::trim())
            .service(statuscode::routes::get_status_code)
    })
    .bind(addr)
    .into_diagnostic()?
    .run()
    .await
    .into_diagnostic()?;
    Ok(())
}
