import express from 'express';
import createHomepageTemplate from './views/index.js';
import createListTemplate from './views/list.js';
import BOOKS_DATA from './data/data.js';
import createBookTemplate from './views/book.js';
import createEditFormTemplate from './views/edit.js';

// create app
const app = express();
app.use(express.urlencoded({extended: false}));

// static assets
app.use(express.static('public'));

// routes
app.get('/', (req, res) => {
  res.send(createHomepageTemplate());
});

// Getting books list
app.get('/books',(req,res) => {
  res.send(createListTemplate());
});

// Adding a book
app.post('/books',(req,res) => {
  const {title, author} = req.body;
  const id = Math.random().toString();
  
  BOOKS_DATA.push({id,title,author});

  res.redirect(`/books/${id}`)
});

// Returns a single book by ID
app.get('/books/:id',(req,res) => {
  const {id} = req.params;
  const book = BOOKS_DATA.find((b) => b.id === id)

  res.send(createBookTemplate(book));
});

// Delete books
app.delete('/books/:id', (req,res) => {
  const id = req.params.id;
  const idx = BOOKS_DATA.findIndex((b) => b.id === id);

  BOOKS_DATA.splice(idx,1);
  // When sending an empty response to HTMX, it will still work normally, but will replace for example with nothing.
  // So can be used for deleting stuff.
  res.send(); 
});

// Updating a book values.
app.put('/books/:id',(req,res) => {
  const {title,author} = req.body;
  const {id} = req.params;

  const newBook = {id,title,author}
  const idx = BOOKS_DATA.findIndex((b) => b.id === id);
   
  BOOKS_DATA[idx] = newBook;
  res.send(createBookTemplate(newBook));
})

// Creating a form to edit a book.
app.get('/books/edit/:id',(req,res) => {
  console.log(req.params)
  const book = BOOKS_DATA.find((b) => b.id === req.params.id);
  res.send(createEditFormTemplate(book))
});

app.post('/books/search',(req,res) => {
  console.log("Invoked")
  const text = req.body.search.toLowerCase();

  const books = BOOKS_DATA.filter((b) => b.title.toLowerCase().includes(text));
  res.send(createListTemplate(books));
})
// listen to port
app.listen(3000, () => {
  console.log('App listening on port 3000');
});