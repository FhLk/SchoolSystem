import React from 'react'
import Error from "../components/Error"
import { Link } from 'react-router-dom'

function ErrorPage() {
    return (
        <div>
            <Link to="/">
                <button class="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
                    Back
                </button>
            </Link>
            <Error />
        </div>
    )
}

export default ErrorPage
