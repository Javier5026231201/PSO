const API = "";

// --- AUTHENTICATION LOGIC ---

async function login() {
    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value.trim();

    if (!username || !password) {
        alert("Username dan password tidak boleh kosong!");
        return;
    }

    try {
        // Mengubah status text button saat loading
        const btn = document.querySelector('button');
        const originalText = btn.innerHTML;
        btn.innerHTML = "Memproses...";
        btn.disabled = true;

        const response = await fetch(`${API}/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ username, password })
        });

        const data = await response.json();

        if (response.ok) {
            // Simulasi menyimpan token session
            localStorage.setItem("isLoggedIn", "true");
            alert(data.message || "Login berhasil!");
            window.location.href = "dashboard.html";
        } else {
            alert(data.message || "Login gagal, periksa kembali kredensial Anda.");
        }
    } catch (error) {
        console.error("Fetch Error:", error);
        alert("Gagal terhubung ke server. Pastikan backend Go sudah berjalan.");
    } finally {
        const btn = document.querySelector('button');
        if(btn) {
            btn.innerHTML = "Login";
            btn.disabled = false;
        }
    }
}

function logout() {
    localStorage.removeItem("isLoggedIn");
    window.location.href = "index.html";
}

function checkAuth() {
    const isLoggedIn = localStorage.getItem("isLoggedIn");
    if (!isLoggedIn) {
        window.location.href = "index.html";
    }
}

// --- NOTES LOGIC ---

async function loadNotes() {
    const notesDiv = document.getElementById("notes");
    if (!notesDiv) return;

    try {
        const response = await fetch(`${API}/notes`);
        if (!response.ok) throw new Error("Gagal mengambil data catatan");
        const notes = await response.json();

        if (notes.length === 0) {
            notesDiv.innerHTML = "<p style='text-align:center; color:#a0aec0;'>Belum ada catatan.</p>";
            return;
        }

        // Hindari innerHTML += di dalam loop (Performa lebih baik)
        const notesHTML = notes.map(note => `
            <div class="data-item">
                <h3>${note.title}</h3>
                <p>${note.content}</p>
                <button onclick="deleteNote(${note.id})" class="delete-btn">Hapus</button>
            </div>
        `).join('');

        notesDiv.innerHTML = notesHTML;
    } catch (error) {
        console.error(error);
        notesDiv.innerHTML = "<p style='color:red;'>Gagal memuat catatan.</p>";
    }
}

async function addNote() {
    const titleInput = document.getElementById("title");
    const contentInput = document.getElementById("content");

    if (!titleInput.value.trim() || !contentInput.value.trim()) {
        alert("Judul dan isi catatan wajib diisi!");
        return;
    }

    try {
        await fetch(`${API}/notes`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                title: titleInput.value,
                content: contentInput.value
            })
        });

        // Kosongkan form setelah submit
        titleInput.value = "";
        contentInput.value = "";
        
        loadNotes();
    } catch(error) {
        alert("Gagal menambahkan catatan");
    }
}

async function deleteNote(id) {
    if (!confirm("Apakah Anda yakin ingin menghapus catatan ini?")) return;

    try {
        await fetch(`${API}/notes/${id}`, {
            method: "DELETE"
        });
        loadNotes();
    } catch(error) {
        alert("Gagal menghapus catatan");
    }
}

// --- SCHEDULES LOGIC ---

async function loadSchedules() {
    const schedulesDiv = document.getElementById("schedules");
    if (!schedulesDiv) return;

    try {
        const response = await fetch(`${API}/schedules`);
        if (!response.ok) throw new Error("Gagal mengambil data jadwal");
        const schedules = await response.json();

        if (schedules.length === 0) {
            schedulesDiv.innerHTML = "<p style='text-align:center; color:#a0aec0;'>Belum ada jadwal.</p>";
            return;
        }

        const schedulesHTML = schedules.map(schedule => `
            <div class="data-item">
                <h3>${schedule.course}</h3>
                <p><strong>Hari:</strong> ${schedule.day} | <strong>Jam:</strong> ${schedule.time}</p>
                <p><strong>Ruang:</strong> ${schedule.room}</p>
                <button onclick="deleteSchedule(${schedule.id})" class="delete-btn">Hapus</button>
            </div>
        `).join('');

        schedulesDiv.innerHTML = schedulesHTML;
    } catch(error) {
        console.error(error);
        schedulesDiv.innerHTML = "<p style='color:red;'>Gagal memuat jadwal.</p>";
    }
}

async function addSchedule() {
    const courseInput = document.getElementById("course");
    const dayInput = document.getElementById("day");
    const timeInput = document.getElementById("time");
    const roomInput = document.getElementById("room");

    if (!courseInput.value.trim() || !dayInput.value.trim()) {
        alert("Mata kuliah dan hari wajib diisi!");
        return;
    }

    try {
        await fetch(`${API}/schedules`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                course: courseInput.value,
                day: dayInput.value,
                time: timeInput.value,
                room: roomInput.value
            })
        });

        // Kosongkan form
        courseInput.value = "";
        dayInput.value = "";
        timeInput.value = "";
        roomInput.value = "";

        loadSchedules();
    } catch(error) {
        alert("Gagal menambahkan jadwal");
    }
}

async function deleteSchedule(id) {
    if (!confirm("Apakah Anda yakin ingin menghapus jadwal ini?")) return;

    try {
        await fetch(`${API}/schedules/${id}`, {
            method: "DELETE"
        });
        loadSchedules();
    } catch(error) {
        alert("Gagal menghapus jadwal");
    }
}