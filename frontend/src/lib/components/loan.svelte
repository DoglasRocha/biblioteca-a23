<script>
	import Bold from './bold.svelte';

	export let loan;

	let loanClass;

	if (new Date() > new Date(loan.return_date)) loanClass = 'bg-danger';
	else if (new Date() < new Date(loan.start_date)) loanClass = 'bg-info';
	else loanClass = '';
</script>

<details class={loanClass}>
	<summary
		>{loan.Request.Book.name} - {loan.Request.Reader.User.name}
		{loan.Request.Reader.User.surname}</summary
	>
	<ul>
		<li>
			<Bold>Livro:</Bold>
			<a href={`/admin/livros/editar/${loan.Request.book_id}`}>
				{loan.Request.Book.name}
			</a>
		</li>
		<li>
			<Bold>Estudante:</Bold>
			<a href={`/admin/leitor/${loan.Request.reader_id}`}>
				{loan.Request.Reader.User.name}
				{loan.Request.Reader.User.surname}
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
			<Bold>Número da cópia do livro:</Bold>
			{loan.Copy.id}
		</li>
		<li>
			<Bold>Renovou:</Bold>
			{loan.has_renewed ? 'Sim' : 'Não'}
		</li>
		<li>
			<Bold>Está atrasado:</Bold>
			{new Date(loan.return_date) < Date.now() ? 'Sim' : 'Não'}
		</li>
	</ul>
	<slot />
</details>
