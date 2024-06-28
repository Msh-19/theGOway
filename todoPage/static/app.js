document.getElementById('add-task').addEventListener('click', function() {
    const taskInput = document.getElementById('task-input');
    const taskName = taskInput.value;

    if (taskName) {
        fetch('/tasks', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ name: taskName })
        })
        .then(response => response.json())
        .then(data => {
            addTaskToList(data);
            taskInput.value = '';
        });
    }
});

function addTaskToList(task) {
    const taskList = document.getElementById('task-list');
    const li = document.createElement('li');
    li.textContent = task.name;
    const deleteButton = document.createElement('button');
    deleteButton.textContent = 'Delete';
    deleteButton.addEventListener('click', function() {
        fetch(`/tasks/${task.id}`, { method: 'DELETE' })
        .then(response => {
            if (response.ok) {
                taskList.removeChild(li);
            }
        });
    });
    li.appendChild(deleteButton);
    taskList.appendChild(li);
}

window.addEventListener('DOMContentLoaded', (event) => {
    fetch('/tasks')
    .then(response => response.json())
    .then(tasks => {
        tasks.forEach(task => {
            addTaskToList(task);
        });
    });
});
