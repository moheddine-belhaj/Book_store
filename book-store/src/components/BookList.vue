<!-- src/components/BookList.vue -->
<template>
    <div>
      <h2>Books</h2>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Author</th>
            <th>Publication</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="book in books" :key="book.ID">
            <td>{{ book.ID }}</td>
            <td>{{ book.name }}</td>
            <td>{{ book.author }}</td>
            <td>{{ book.publication }}</td>
            <td>
              <button @click="confirmDelete(book.ID)">Delete</button>
              <button @click="editBook(book)">Edit</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  <script>
  import BookService from '../services/BookService';
  
  export default {
    props: {
      books: Array
    },
    methods: {
      confirmDelete(bookID) {
        const confirmed = window.confirm('Are you sure you want to delete this book?');
        if (confirmed) {
          this.deleteBook(bookID);
        }
      },
  
      async deleteBook(bookID) {
        if (!bookID) {
          console.error('Error: bookID is undefined or null');
          return;
        }
        try {
          await BookService.deleteBook(bookID);
          this.$emit('book-deleted');
        } catch (error) {
          console.error('Error deleting book:', error);
        }
      },
  
      editBook(book) {
        this.$emit('edit-book', book);
      }
    }
  };
  </script>
  