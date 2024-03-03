<script>
	import Card from '$lib/components/card.svelte';
	import FormField from '$lib/components/form-field.svelte';
	import BlankForm from '$lib/components/blank-form.svelte';
	import axios from 'axios';
	import { api } from '$lib/utils/login.js';

	let book = {
		name: '',
		isbn: '',
		description: '',
		gender: ''
	};

	let errorFromServer;

	const handleISBNSubmit = async () => {
		// 8575421131 - codigo para teste
		// usando openlibray: `https://openlibrary.org/isbn/${book.isbn}.json`
		// usando google books api: `https://www.googleapis.com/books/v1/volumes?q=isbn:${book.isbn}`
		let bookData = await axios.get(
			`https://www.googleapis.com/books/v1/volumes?q=isbn:${book.isbn}`
		);

		if (bookData.data.totalItems > 0) {
			book.name = bookData.data.items[0].volumeInfo.title ?? '';
			book.description = bookData.data.items[0].volumeInfo.description ?? '';
			book = book;
		}
	};

	const handleSubmit = async () => {
		try {
			let request = await api.post('/admin/livros/cadastrar', book);

			if (request.status == 201) {
				document.location.href = '/admin/livros/buscar';
				return;
			}
		} catch (error) {
			errorFromServer = error.response.data;
		}
	};
</script>

<Card class="w-75">
	<h1 class="text-center">Cadastro de livro</h1>
	<BlankForm class="container">
		<div class="row">
			<div class="col">
				<FormField name="book-name" label="Nome do livro" bind:value={book.name} required />
			</div>
			<div class="col">
				<div class="d-flex align-items-end">
					<FormField name="book-isbn" label="ISBN" type="number" bind:value={book.isbn} />
					<button
						class="ms-1 btn btn-outline-success button"
						type="button"
						on:click={handleISBNSubmit}>Pesquisar</button
					>
				</div>
			</div>
		</div>
		<div class="row">
			<div class="col">
				<div class="mt-3 d-flex flex-column">
					<label for="book-isbn" class="form-label">Descrição do livro</label>
					<textarea
						class="form-control"
						name="book-isbn"
						id="book-isbn"
						rows="5"
						bind:value={book.description}
					/>
				</div>
			</div>
			<div class="col">
				<FormField name="book-gender" label="Gênero do livro" bind:value={book.gender} />
			</div>
			{#if errorFromServer}
				<div class="mt-3 text-danger">
					<p>Ocorreu um erro. Talvez isso ajude:</p>
					<p>{errorFromServer}</p>
				</div>
			{/if}
			<div class="d-flex justify-content-end mt-3">
				<button type="submit" class="btn btn-primary" on:click={handleSubmit}>Cadastrar</button>
			</div>
		</div>
	</BlankForm>
</Card>

<style>
	.button {
		height: 50%;
	}
</style>
