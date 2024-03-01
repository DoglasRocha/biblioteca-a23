<script>
	import Bold from './bold.svelte';

	export let loan;
</script>

<details class={new Date() > new Date(loan.return_date) && !loan.has_returned ? 'bg-danger' : ''}>
	<summary>
		{loan.Request.Book.name}
	</summary>
	<ul>
		<li>
			<Bold>Livro:</Bold>
			<a href={`/emprestar/${loan.Request.book_id}`}>
				{loan.Request.Book.name}
			</a>
		</li>
		<li>
			<Bold>Data de empréstimo:</Bold>
			{new Date(loan.start_date).toLocaleDateString('pt-br')}
		</li>
		<li>
			<Bold>Data de devolução:</Bold>
			{new Date(loan.return_date).toLocaleDateString('pt-br')}
		</li>
		<li>
			<Bold>Renovou:</Bold>
			{loan.has_renewed ? 'Sim' : 'Não'}
		</li>
		<li>
			<Bold>Está atrasado:</Bold>
			{new Date(loan.return_date) < Date.now() && !loan.has_returned ? 'Sim' : 'Não'}
		</li>
	</ul>
	<slot />
</details>
