use actix_web::{get, HttpRequest, HttpResponse, Responder};
use http::StatusCode;

#[get("/{status_code}")]
pub async fn get_status_code(req: HttpRequest) -> impl Responder {
    if let Some(status_code) = req.match_info().get("status_code") {
        let status_code = status_code.parse::<StatusCode>();
        if let Ok(status_code) = status_code {
            HttpResponse::build(status_code).body(status_code.canonical_reason().unwrap_or(""))
        } else {
            HttpResponse::NotFound().body("Invalid status code")
        }
    } else {
        HttpResponse::NotFound().body("Invalid status code")
    }
}
