<script>
	import Bold from './bold.svelte';
	import { api } from '../utils/api';
	export let request;

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

	let startDate = getNextSaturday(),
		returnDate = sum6Days(startDate);
</script>

<details>
	<summary>{request.Book.name}</summary>
	<ul>
		<li>
			<Bold>Livro:</Bold>
			<a href={`/emprestar/${request.book_id}`}>{request.Book.name}</a>
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
</details>
