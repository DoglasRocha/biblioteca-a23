<script>
	import Card from '$lib/components/card.svelte';
	import Link from '$lib/components/link.svelte';
</script>

<Card class="max-width-60">
	<h1 class="text-center">Biblioteca do Cursinho A23</h1>
	<div class="d-flex justify-content-center">
		<img src="/logo.png" alt="Logo do Cursinho A23" class="logo" />
	</div>
	<p class="mt-3">
		Seja bem-vind@ à biblioteca do Cursinho Popular Alicerce 23! Aqui você pode emprestar um dos
		livros do nosso acervo.
	</p>
	<p>
		Para procurar um livro para emprestar, clique em <Link href="/buscar">Buscar Livro</Link> na barra
		de navegação.
	</p>
	<p>
		Para verificar seu histórico de empréstimos, clique em <Link href="/historico">Histórico</Link> na
		barra de navegação.
	</p>
	<p>
		Para renovar o livro que você está emprestando atualmente, clique em <Link href="/renovar"
			>Renovar</Link
		> na barra de navegação.
	</p>
</Card>

<style>
	:global(.max-width-60) {
		width: 60%;
	}

	.logo {
		width: 200px;
	}
</style>
