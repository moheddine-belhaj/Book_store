<template>
  <div id="app">
    <h1>Book Store</h1>

    <!-- Create Book Form -->
    <div>
      <h2>Create a New Book</h2>
      <form @submit.prevent="createBook">
        <input v-model="newBook.name" placeholder="Book Name" required />
        <input v-model="newBook.author" placeholder="Author" required />
        <input v-model="newBook.publication" placeholder="Publication" required />
        <button type="submit" :disabled="isCreating">Add Book</button>
      </form>
    </div>

    <!-- Book List -->
    <div>
      <h2>Books</h2>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Author</th>
            <th>Publication</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="book in books" :key="book.id">
            <td>{{ book.name }}</td>
            <td>{{ book.author }}</td>
            <td>{{ book.publication }}</td>
            <td>
              <button @click="deleteBook(book.id)">Delete</button>
              <button @click="editBook(book)">Edit</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Edit Book Form -->
    <div v-if="editingBook">
      <h2>Edit Book</h2>
      <form @submit.prevent="updateBook">
        <input v-model="editingBook.name" placeholder="Book Name" required />
        <input v-model="editingBook.author" placeholder="Author" required />
        <input v-model="editingBook.publication" placeholder="Publication" required />
        <button type="submit" :disabled="isUpdating">Update Book</button>
      </form>
    </div>
  </div>
</template>

<script>
import BookService from './services/BookService';

export default {
  data() {
    return {
      books: [],
      newBook: {
        name: '',
        author: '',
        publication: ''
      },
      editingBook: null,
      isCreating: false,
      isUpdating: false
    };
  },

  created() {
    this.fetchBooks();
  },

  methods: {
    async fetchBooks() {
      try {
        const response = await BookService.getAllBooks();
        this.books = response.data;
      } catch (error) {
        console.error('Error fetching books:', error);
      }
    },

    async createBook() {
      this.isCreating = true;
      try {
        await BookService.createBook(this.newBook);
        this.fetchBooks();
        this.newBook = { name: '', author: '', publication: '' };
      } catch (error) {
        console.error('Error creating book:', error);
      } finally {
        this.isCreating = false;
      }
    },

    async deleteBook(bookId) {
      try {
        await BookService.deleteBook(bookId);
        this.fetchBooks();
      } catch (error) {
        console.error('Error deleting book:', error);
      }
    },

    editBook(book) {
      this.editingBook = { ...book };
    },

    async updateBook() {
      this.isUpdating = true;
      try {
        await BookService.updateBook(this.editingBook.id, this.editingBook);
        this.fetchBooks();
        this.editingBook = null;
      } catch (error) {
        console.error('Error updating book:', error);
      } finally {
        this.isUpdating = false;
      }
    }
  }
};
</script>

<style scoped>
#app {
  font-family: Arial, sans-serif;
  margin: 20px;
}

form {
  margin-bottom: 20px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  padding: 10px;
  border: 1px solid #ddd;
}

button {
  margin-right: 10px;
}

button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
</style>
