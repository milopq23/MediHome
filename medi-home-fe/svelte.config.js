// import adapter from '@sveltejs/adapter-auto';
import adapter from '@sveltejs/adapter-static';

const dev = process.env.NODE_ENV === 'development';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		// adapter-auto only supports some environments, see https://svelte.dev/docs/kit/adapter-auto for a list.
		// If your environment is not supported, or you settled on a specific environment, switch out the adapter.
		// See https://svelte.dev/docs/kit/adapters for more information about adapters.
		adapter: adapter(
			{
				pages: 'build',
				assets: 'build',
				fallback: 'index.html', // hoặc 'index.html' nếu dùng SPA routing
				precompress: false
			}),
		paths:{
			base: dev?'':'/MediHome'
		}		
	},
};

export default config;
