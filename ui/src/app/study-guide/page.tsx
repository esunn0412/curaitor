"use client"
import './study-guide.css'
import {createRoot} from 'react-dom/client'
import Markdown from 'react-markdown'

const StudyGuide=() => {
    return <Markdown>{markdown}</Markdown>
}

const markdown = `
# Lin Alg 
## Section 1.1 
 Blah Blah Blah Blah
 ## Section 1.2
 Blah Bak Bkha
 `;


// createRoot(document.body).render(<Markdown>{markdown}</Markdown>)
export default StudyGuide