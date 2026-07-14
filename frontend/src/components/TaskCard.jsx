import { deleteTask } from "../api";

function TaskCard({ task, onDeleted }) {

    async function handleDelete() {

        if (!window.confirm(`Delete "${task.title}"?`)) {
            return;
        }

        try {

            await deleteTask(task.id);

            onDeleted();

        } catch (err) {

            alert("Cannot delete task");

        }

    }

    return (

        <div className="task-card">

            <div className="task-header">

                <h3>{task.title}</h3>

                <span className={`priority ${task.priority}`}>
                    {task.priority}
                </span>

            </div>

            <p>

                {task.description || "No description"}

            </p>

            <div className="task-footer">

                <span>

                    Status: {task.status}

                </span>

                <div className="buttons">

                    <button>

                        Edit

                    </button>

                    <button
                        className="delete"
                        onClick={handleDelete}
                    >

                        Delete

                    </button>

                </div>

            </div>

        </div>

    );

}

export default TaskCard;