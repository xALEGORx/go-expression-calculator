import React, { Component } from "react";
import Task from "./task";
import ApiService from "./../../api";

export default class Tasks extends Component {
    constructor(props) {
        super(props);

        this.state = {
            tasks: []
        };
    }

    componentDidMount() {
        this.setState({ content: [] })
        ApiService.getTasks().then(
            response => {
                this.setState({tasks: response.data.data})
            }
        )
    }

    render() {
        return (
            <div>
                <div className="w-1/2 mb-5">
                    <h1 className="text-black text-xl mb-1">Добавить новую задачу</h1>
                    <input className="border border-gray w-1/2 rounded-md p-1 mr-1" placeholder="2+2"></input>
                    <button className="bg-blue p-1 rounded-md px-5">Решить</button>
                </div>
                <div>
                    <h1 className="text-black text-xl mb-1">Список всех задач</h1>
                    <div className="flex flex-col gap-2">
                        {this.state.tasks.map(task => {
                            console.log(task)
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
