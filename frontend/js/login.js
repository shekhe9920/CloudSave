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


const loginForm = document.getElementById("loginForm");
const registerForm = document.getElementById("registerForm")

document.getElementById("showRegister").addEventListener("click", (e) => {
    e.preventDefault();
    loginForm.style.display = "none";
    registerForm.style.display = "flex"
})

document.getElementById("showLogin").addEventListener("click", (e) => {
    e.preventDefault();
    loginForm.style.display = "none";
    registerForm.style.display = "flex"
})

registerForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const email = document.getElementById("regEmail").values;
    const password = document.getElementById("regPassword").value;

    try {
        const response = await fetch("http://localhost:8080/regoister", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({ email, password })
        });

        if (!response.ok) {
            throw new Error("Registration failed");
        }

        alert("Account created! You can now log in.");

        registerForm.reset()
        registerForm.style.display = "none";
        loginForm.display = "flex";

    } catch (err) {
        alert(err.message);
    }
    
})

























