import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { LoadWasm } from './LoadWasm/index.tsx'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <LoadWasm>
      <App/>
    </LoadWasm>
  </React.StrictMode>,
)
