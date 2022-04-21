use actix_web::{get, web::Path, HttpResponse, Responder};
use http::StatusCode;

#[get("/{status_code}")]
pub async fn get_status_code(path: Path<u16>) -> impl Responder {
    let status_code = StatusCode::from_u16(path.into_inner());
    match status_code {
        Ok(status_code) => {
            HttpResponse::build(status_code).body(status_code.canonical_reason().unwrap_or(""))
        }
        Err(_) => HttpResponse::NotFound().body("Invalid status code"),
    }
}
