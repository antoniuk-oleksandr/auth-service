document.addEventListener("DOMContentLoaded", () => {
  document.body.addEventListener("login:success", (e) => {
    localStorage.setItem("access_token", e.detail.access_token);
    alert("Login successful!");
  });

  document.body.addEventListener("register:success", (e) => {
    localStorage.setItem("access_token", e.detail.access_token);
    alert("Registration successful!");
  });

  document.body.addEventListener("login:error", (e) => {
    alert("Login failed: " + e.detail.error);
  });

  document.body.addEventListener("register:error", (e) => {
    alert("Registration failed: " + e.detail.message);
  });
});
