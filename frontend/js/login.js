// login.js

/**
 * Handles login form submission.
 *
 * - Prevents default form submission (page reload)
 * - Collects user credentials from input fields
 * - Sends credentials to the backend login endpoint using fetch
 * - Redirects the user to the dashboard on successful login
 */
document.getElementById("loginForm").addEventListener("submit", async (e) => {
    // Prevent the browser's default form submission behavior
    e.preventDefault();

    // Retrieve user input values
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    try {
        // Send login request to backend
        const response = await fetch("http://localhost:8080/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email: email,
                password: password
            })
        });

        // Check if authentication failed
        if (!response.ok) {
            throw new Error("Invalid login credentials");
        }

        // Parse JSON response from backend
        const data = await response.json();

        // authentication token
        localStorage.setItem("token", data.token);

        // Redirect user to dashboard after successful login
        window.location.href = "dashboard.html";

    } catch (err) {
        alert(err.message);
    }
});
