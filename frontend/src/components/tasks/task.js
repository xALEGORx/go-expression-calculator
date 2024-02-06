import React, { useEffect } from 'react'

export default function Task({ task }) {
    return (
        <div className="border border-gray flex items-center rounded-md p-3 justify-between">
            <div className="flex items-center gap-5">
                <span className={`w-8 h-8 inline-block rounded-full bg-${task.status}`}></span>
                <span className="text-black font-semibold text-xl">{task.expression}</span>
                <span className={`text-gray`}>{prepareStatus(task.status)}</span>
            </div>
            <div>
                <span className="text-gray text-sm">07.02.2024 00:21:27</span>
            </div>
        </div>
    )
}

function prepareStatus(status) {
    switch (status) {
        case "created":
            return "В процессе"
        default:
            return "Unknown status"
    }
}