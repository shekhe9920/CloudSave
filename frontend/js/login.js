document.getElementById("loginForm").addEventListener("submit", async (e) => {
    e.preventDefault(); // stopper vanlig form-submit

    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    try {
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

        if (!response.ok) {
            throw new Error("Invalid login credentials");
        }

        const data = await response.json();

        alert("Login successful!");

        // TODO:
        // localStorage.setItem("token", data.token);

        window.location.href = "dashboard.html";


    } catch (err) {
        alert(err.message);
    }
});
