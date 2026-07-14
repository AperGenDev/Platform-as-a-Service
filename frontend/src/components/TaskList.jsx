import TaskCard from "./TaskCard";

function TaskList({ tasks, onDeleted }) {

    return (

        <section className="card">

            <h2>

                Tasks ({tasks.length})

            </h2>

            {

                tasks.length === 0

                    ?

                    <p>No tasks</p>

                    :

                    tasks.map(task => (

                        <TaskCard
                            key={task.id}
                            task={task}
                            onDeleted={onDeleted}
                        />

                    ))

            }

        </section>

    );

}

export default TaskList;