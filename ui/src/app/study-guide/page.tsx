import './study-guide.css'
import Markdown from 'react-markdown'
import QuizButton from './quizbutton';

const markdown = `
# Lin Alg 
## Section 1.1 
 Blah Blah Blah Blah
## Section 1.2
 Blah Bak Bkha
 `;


const StudyGuide=() => {
    return (   
    <>
        <QuizButton />
        
        <Markdown>{markdown}</Markdown>
    </>
    )
}


export default StudyGuide