import React, { Component } from "react";
import Task from "./task";
import ApiService from "./../../api";

export default class Tasks extends Component {
    constructor(props) {
        super(props);

        this.state = {
            tasks: [],
            expression: "",
        };

        this.handleAddTask = this.handleAddTask.bind(this);
        this.getTaskList = this.getTaskList.bind(this);
    }

    handleExpressionChange = (event) => {
        this.setState({ expression: event.target.value });
    };

    handleAddTask() {
        this.setState({ expression: "" })

        ApiService.addTask(this.state.expression).then((response) => {
            this.getTaskList()
        })
    }

    getTaskList() {
        ApiService.getTasks().then(
            response => {
                this.setState({ tasks: response.data.data })
            }
        )
    }

    componentDidMount() {
        this.getTaskList()

        this.interval = setInterval(() => {
            this.getTaskList()
        }, 1000);
    }

    componentWillUnmount() {
      clearInterval(this.interval);
    }

    render() {
        return (
            <div>
                <div className="w-1/2 mb-5">
                    <h1 className="text-black text-xl mb-1">Добавить новую задачу</h1>
                    <input className="border border-gray w-1/2 rounded-md p-1 mr-1 text-black" placeholder="Введите пример" onChange={this.handleExpressionChange} value={this.state.expression}></input>
                    <button className="bg-blue p-1 rounded-md px-5" onClick={this.handleAddTask}>Решить</button>
                </div>
                <div>
                    <h1 className="text-black text-xl mb-1">Список всех задач</h1>
                    <div className="flex flex-col gap-2">
                        {this.state.tasks.map(task => {
                            return (
                                <Task key={task.task_id} task={task} />
                            )
                        })}
                    </div>
                </div>
            </div>
        );
    }
}
