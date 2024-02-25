<script>
	import Bold from './bold.svelte';
	import { api } from '../utils/api';
	export let request;
	let error;

	const getNextSaturday = () => {
		let weekday = new Date().getDay(),
			nextSaturday = new Date();

		nextSaturday.setDate(nextSaturday.getDate() + (weekday == 6 ? 7 : 6 - weekday));

		return nextSaturday;
	};

	const sum6Days = (date) => {
		let newDate = new Date(date.getTime());
		newDate.setDate(newDate.getDate() + 7);

		return newDate;
	};

	const handleOrder = async (order, request_id) => {
		try {
			let request =
				order == 'aprovar'
					? await api.post(`/admin/emprestimos/${order}/${request_id}`)
					: await api.delete(`/admin/emprestimos/${order}/${request_id}`);

			if (request.status == 200) document.location.href += '/';
		} catch (err) {
			error = err.response.data;
		}
	};

	let startDate = getNextSaturday(),
		returnDate = sum6Days(startDate);
</script>

<details>
	<summary>{request.Book.name} - {request.Reader.User.name} {request.Reader.User.surname}</summary>
	<ul>
		<li>
			<Bold>Estudante:</Bold>
			<a href={`/admin/usuario/${request.Reader.user_id}`}>
				{request.Reader.User.name}
				{request.Reader.User.surname}
			</a>
		</li>
		<li>
			<Bold>Livro:</Bold>
			<a href={`/admin/livros/editar/${request.book_id}`}>{request.Book.name}</a>
		</li>
		<li>
			<Bold>Data de início:</Bold>
			{startDate.toLocaleDateString('pt-br')}
		</li>
		<li>
			<Bold>Data provável de retorno:</Bold>
			{returnDate.toLocaleDateString('pt-br')}
		</li>
	</ul>
	<div class="d-flex justify-content-end">
		<button
			type="button"
			class="btn btn-danger ms-1"
			on:click={() => handleOrder('rejeitar', request.id)}
		>
			Rejeitar
		</button>
		<button
			type="button"
			class="btn btn-success ms-1"
			on:click={() => handleOrder('aprovar', request.id)}
		>
			Aprovar
		</button>
	</div>
	{#if error}
		<div class="text-danger">
			<p>Ops, aconteceu um erro. Talvez isso ajude:</p>
			<p>{error}</p>
		</div>
	{/if}
</details>
