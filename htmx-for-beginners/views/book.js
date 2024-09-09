const createBookTemplate = (book) => /*html*/`
    <li data-id="${book.id}">
        <div 
        class="details"
        hx-get="/books/edit/${book.id}"
        hx-target="closest li"
        >
            <h3>${book.title}</h3>
            <p>${book.author}</p>
        </div>
        <!-- This (hx-target="closest li") will go up the tree, to the first li and replace it, meaning this whole element -->
        <button
         hx-delete="/books/${book.id}" 
         hx-target="closest li" 
         hx-swap="outerHTML"
         >Delete</button>
    </li>
`;

export default createBookTemplate;