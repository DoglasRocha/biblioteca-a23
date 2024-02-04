<script>
	import Card from '$lib/components/card.svelte';
	import FormField from '$lib/components/form-field.svelte';
	import axios from 'axios';

	let name = '';
	let book = {
		name: '',
		isbn: '',
		description: '',
		gender: ''
	};

	const handleISBNSubmit = async () => {
		let bookData = await axios.get(`https://openlibrary.org/isbn/${book.isbn}.json`);
		name = bookData.data.title;
		book.description = bookData.data.description;
		book = book;
		console.log(book);
		console.log(name);
	};
</script>

<Card class="w-75">
	<h1 class="text-center">Cadastro de livro</h1>
	<div class="container">
		<div class="row">
			<div class="col">
				<FormField name="book-name" label="Nome do livro" bind:value={name} required />
			</div>
			<div class="col">
				<div class="d-flex align-items-end">
					<FormField name="book-isbn" label="ISBN" type="number" bind:value={book.isbn} />
					<button class="ms-1 btn btn-outline-success button" on:click={handleISBNSubmit}
						>Pesquisar</button
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
				<FormField name="book-gender" label="Gênero do livro" bind:value={book.gender} required />
			</div>
		</div>
	</div>
</Card>

<style>
	.button {
		height: 50%;
	}
</style>
