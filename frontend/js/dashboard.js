// dashboard.js

const DEV_MODE = true; // 👈 sett til false i produksjon

document.addEventListener("DOMContentLoaded", () => {

    /* =========================
       AUTH CHECK
    ========================= */
    if (!DEV_MODE) {
        const token = localStorage.getItem("token");
        if (!token) {
            window.location.href = "login.html";
            return;
        }
    }

    /* =========================
       STORAGE PIE CHART
    ========================= */
    const canvas = document.getElementById("storageChart");

    // Sikkerhet: sjekk at canvas finnes
    if (!canvas) {
        console.warn("storageChart canvas not found");
        return;
    }

    const ctx = canvas.getContext("2d");

    // Fake data (senere fra backend)
    const data = [
        { label: "Documents", value: 2, color: "#4e79a7" },
        { label: "Images", value: 3, color: "#59a14f" },
        { label: "Videos", value: 4, color: "#f28e2b" },
        { label: "Other", value: 1, color: "#e15759" }
    ];

    const total = data.reduce((sum, item) => sum + item.value, 0);

    let startAngle = 0;

    data.forEach(item => {
        const sliceAngle = (item.value / total) * 2 * Math.PI;

        ctx.beginPath();
        ctx.moveTo(90, 90);
        ctx.arc(90, 90, 80, startAngle, startAngle + sliceAngle);
        ctx.closePath();

        ctx.fillStyle = item.color;
        ctx.fill();

        startAngle += sliceAngle;
    });

});
