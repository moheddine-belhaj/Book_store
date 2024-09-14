<template>
    <div>
      <h2 v-if="!editingBook">Create a New Book</h2>
      <h2 v-else>Edit Book</h2>
      <form @submit.prevent="handleSubmit">
        <input v-model="localBook.name" placeholder="Book Name" required />
        <input v-model="localBook.author" placeholder="Author" required />
        <input v-model="localBook.publication" placeholder="Publication" required />
        <button type="submit" :disabled="isProcessing">{{ buttonText }}</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      book: {
        type: Object,
        default: () => null // Ensure book has a default value
      },
      isProcessing: Boolean
    },
    data() {
      return {
        localBook: { ...this.book } // Create a local copy of the book prop
      };
    },
    computed: {
      buttonText() {
        return this.editingBook ? 'Update Book' : 'Add Book';
      },
      editingBook() {
        // Check if book is not null and has an ID property
        return this.book && this.book.ID;
      }
    },
    methods: {
      async handleSubmit() {
        if (this.editingBook) {
          this.$emit('update-book', this.localBook);
        } else {
          this.$emit('create-book', this.localBook);
        }
      }
    },
    watch: {
      book(newBook) {
        this.localBook = { ...newBook }; // Update localBook when the book prop changes
      }
    }
  };
  </script>
  