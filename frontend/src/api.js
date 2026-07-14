const API_URL = "/api";

export async function checkHealth() {
    const response = await fetch(`${API_URL}/health`);

    if (!response.ok) {
        throw new Error("API unavailable");
    }

    return response.json();
}

export async function getTasks() {
    const response = await fetch(`${API_URL}/tasks`);

    if (!response.ok) {
        throw new Error("Cannot load tasks");
    }

    return response.json();
}

export async function createTask(task) {
    const response = await fetch(`${API_URL}/tasks`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(task),
    });

    if (!response.ok) {
        throw new Error("Cannot create task");
    }

    return response.json();
}

export async function deleteTask(id) {
    const response = await fetch(`${API_URL}/tasks/${id}`, {
        method: "DELETE",
    });

    if (!response.ok) {
        throw new Error("Cannot delete task");
    }
}

export async function updateTask(id, updates) {
    const response = await fetch(`${API_URL}/tasks/${id}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(updates),
    });

    if (!response.ok) {
        throw new Error("Cannot update task");
    }

    return response.json();
}