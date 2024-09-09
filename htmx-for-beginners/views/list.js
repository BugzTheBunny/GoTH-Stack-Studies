import createBookTemplate from "./book.js";

const createListTemplate = (books) => /*html*/`
    <ul>
        <li>
            ${books.map((book) => createBookTemplate(book)).join('')}
        </li>
    </ul>
`;

export default createListTemplate;