import React from 'react'

export default function Task({ task }) {
    const statuscoloros = {
        "completed": "bg-[#08c708]",
        "created": "bg-[#fbff03]",
        "processed": "bg-[#ffae00]",
        "fail": "bg-[#c70808]"
    };

    const taskColor = statuscoloros[task.status];

    return (
        <div className="border border-gray flex items-center rounded-md p-3 justify-between">
            <div className="flex items-center gap-5">
                <span className={`w-8 h-8 inline-block rounded-full ${taskColor}`}></span>
                <span className="text-black font-semibold text-xl">{task.expression}{task.status === "completed" ? " = " + task.answer : ""}</span>
                <span className={`text-gray`}>{prepareStatus(task)}</span>
                {task.status === "fail" ? (
                    <span className="text-[#c70808] font-semibold text-xl">
                        {task.answer}
                    </span>
                ) : (<span></span>)}
            </div>
            <div>
                <span className="text-completed"></span>
                <span className="text-gray text-sm">07.02.2024 00:21:27</span>
            </div>
        </div>
    )
}

function prepareStatus(task) {
    switch (task.status) {
        case "created":
            return "Создан"
        case "processed":
            return `В процессе (Агент ${task.agent_id})`
        case "completed":
            return "Завершен"
        case "fail":
            return "Ошибка"
        default:
            return "Unknown status"
    }
}