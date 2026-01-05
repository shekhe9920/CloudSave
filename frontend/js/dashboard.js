// dashboard.js

const DEV_MODE = true; // 👈 sett til false i produksjon

if (!DEV_MODE) {
    const token = localStorage.getItem("token");
    if (!token) {
        window.location.href = "index.html";
    }
}


// Check if user has token in localStorage
const token = localStorage.getItem("token");

if (!token) {
    window.location.href = "login.html";
} else {
    document.getElementById("dashboardContent").innerHTML = "<h2>Welcome to the Dashboard</h2>";

    // Add a logout event listener to the button
    document.getElementById("logoutBtn").addEventListener("click", () => {
        localStorage.removeItem("token");
        window.location.href = "index.html";  // Redirect to login after logout
    });
}
