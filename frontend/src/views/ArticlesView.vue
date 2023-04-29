<script setup>
</script>

<template>
    <div class="list-articles md:h-full h-[110%]">
        <h1 class="md:ps-5 text-5xl font-bold">Новости
<!--            <% if user_signed_in? %>-->
<!--            <span><a href="/articles/new">+</a></span>-->
<!--            <% end %>-->
        </h1>
        <div class="grid md:grid-cols-2">
<!--            <% @articles.each do |a| %>-->
<!--            &lt;!&ndash;      <a href="/articles/<%#= a.id %>" class="block mt-3 p-1 md:basis-1/3 basis-1/2 ">&ndash;&gt;-->
<!--            <%#= image_tag a.image, class:"h-56 w-full rounded-bl-3xl rounded-tr-3xl object-cover sm:h-64 lg:h-72" %>-->

            <!--        <div class="mt-4 sm:flex text-3xl sm:items-center sm:justify-center sm:gap-4">-->
            <!--          <strong class="font-medium text-center"><%#= a.title %></strong>-->

            <!--          <span class="hidden sm:block sm:h-px sm:w-8 sm:bg-yellow-500"></span>-->

            <!--          <p class="mt-0.5 opacity-50 sm:mt-0"><%#= User.find_by(id:a.author_id).login %></p>-->
            <!--        </div>-->
            <!--      </a>-->
            <article class="flex bg-white m-2 transition hover:shadow-xl">
                <div class="rotate-180 p-2 [writing-mode:_vertical-lr]">
                    <time datetime="2022-10-10" class="flex items-center justify-between gap-4 text-xs font-bold uppercase text-gray-900">
                        <span>2022</span>
                        <span class="w-px flex-1 bg-gray-900/10"></span>
                        <span>Oct 10</span>
                    </time>
                </div>

                <div class="hidden sm:block sm:basis-56">
<!--                    <%= image_tag a.image, class:"aspect-square h-full w-full object-cover" %>-->

                </div>

                <div class="md:bg-white bg-[url('')] flex flex-1 flex-col justify-between">
                    <div class="border-l border-gray-900/10 p-4 sm:border-l-transparent sm:p-6">
                        <a href="/articles/<%= a.id %>">
                            <h3 class="font-bold uppercase text-gray-900">
<!--                                <%= a.title %>-->
                            </h3>
                        </a>

                        <p class="mt-2 line-clamp-3 text-sm/relaxed text-gray-700">
<!--                            <%= a.body.slice(0, 250) %>-->
                        </p>
                    </div>

                    <div class="sm:flex sm:items-end sm:justify-end">
                        <a href="#" class="block bg-yellow-300 px-5 py-3 text-center text-xs font-bold uppercase text-gray-900 transition hover:bg-yellow-400">
                            Read Blog
                        </a>
                    </div>
                </div>
            </article>

<!--            <% end %>-->
        </div>
    </div>

</template>


<style scoped>
.list-articles {
    overflow-y:scroll;
}
.list-articles {
    -ms-overflow-style: none;  /* Internet Explorer 10+ */
    scrollbar-width: none;  /* Firefox */
}
.list-articles::-webkit-scrollbar {
    display: none;  /* Safari and Chrome */
}
</style>