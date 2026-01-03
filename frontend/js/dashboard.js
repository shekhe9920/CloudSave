// dashboard.js

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
