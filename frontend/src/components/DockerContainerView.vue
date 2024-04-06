<script setup>
import {onServerPrefetch, ref} from 'vue';
import {useMutation, useQuery, useQueryClient} from '@tanstack/vue-query'
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Tag from 'primevue/tag';
import Button from 'primevue/tag';
import {formatDistance, fromUnixTime} from "date-fns";
import 'primeicons/primeicons.css'
import InputText from 'primevue/inputtext';
import Menu from 'primevue/menu';
import Dialog from 'primevue/dialog';
import ContainerDetail from "./ContainerDetail.vue";

let imageInputText = ref('');
let imageName = ref('');
let containerInputText = ref('');
let containerName = ref('');
let timer = null;
const menu = ref();
const selectedContainerId = ref('');
const selectedContainer = ref({});
let visible = ref(false);

const items = ref([
  {
    items: [
      {
        label: 'Restart',
        icon: 'pi pi-refresh',
        command: ()=> {
          restartContainer(selectedContainerId.value)
        }
      },
      {
        label: 'Stop',
        icon: 'pi pi-history',
        command: ()=> stopContainer(selectedContainerId.value)
      },
      {
        label: 'Remove',
        icon: 'pi pi-trash',
        command: ()=> removeContainer(selectedContainerId.value)
      }
    ]
  }
]);

const queryClient = useQueryClient()

/*onServerPrefetch(async () => {

})*/

const {isFetching, isError, data: containers, error} = useQuery({
  queryKey: ['containers', imageName, containerName],
  queryFn: async () => await getInfo(imageName.value, containerName.value)
})


const {mutate: restartContainer} = useMutation({
  mutationFn: (id) => {
    mutationFetcher(`/server/containers/${id}/restart`)
        .then(() => {
          queryClient.invalidateQueries({queryKey: ['containers']})
        })
  }
})
const {mutate: stopContainer} = useMutation({
  mutationFn: (id) => {
    mutationFetcher(`/server/containers/${id}/stop`)
        .then(() => {
          queryClient.invalidateQueries({queryKey: ['containers']})
        })
  },
})
const {mutate: removeContainer} = useMutation({
  mutationFn: (id) => {
    mutationFetcher(`/server/containers/${id}`, "DELETE")
        .then(() => {
          queryClient.invalidateQueries({queryKey: ['containers']})
        })
  },
})

const mutationFetcher = async (url, method = "POST") => {
  return await fetch(url, {
    method,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

const getInfo = async (imageName, containerName) => {
  const response = await fetch(`/server/containers/json?imageName=${imageName}&containerName=${containerName}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  });
  const data = await response.json()
  return data.data;
}

const getSeverity = (container) => {
  switch (container.State) {
    case 'created':
      return 'success';

    case 'running':
      return 'success';

    case 'exited':
      return 'danger';

    case 'removing':
      return 'warning';

    case 'paused':
      return 'info';

    case 'dead':
      return 'danger';

    case 'restarting':
      return 'warning';

    default:
      return null;
  }
};

const formatDate = (data) => {
  return formatDistance(new Date(fromUnixTime(data)), new Date())
}

const getContainerName = (names) => {
  return names[0].slice(1)
}

const handleImageChange = (e) => {
  if (timer) clearTimeout(timer);
  imageInputText = e.target.value;
  timer = setTimeout(() => {
    imageName.value = e.target.value
  }, 1000);
}

const handleContainerChange = (e) => {
  if (timer) clearTimeout(timer);
  containerInputText = e.target.value;
  timer = setTimeout(() => {
    containerName.value = e.target.value
  }, 1000);
}

const getShortId = (id) => {
  return id.slice(0, 12)
}

const toggle = (event) => {
  menu.value.toggle(event);
};

const openModal = (container) => {
  selectedContainer.value = container
  visible.value = true
}


</script>
<template>

  <div>
    <div class="w-full h-16 bg-cyan-700 flex items-center">
      <h1 class="font-bold text-white text-lg px-6">Kinexon docker runtime</h1>
    </div>

    <div class="shadow-lg m-4 rounded border">
      <DataTable :value="containers" tableStyle="min-width: 30rem" :loading="isFetching">
        <template #header>
          <div class="flex items-center justify-between">
            <div class="flex items-center justify-center">
              <span class="text-xl text-900 font-bold mr-3">Containers</span>
              <span class="inline px-2 py-1 rounded-lg bg-gray-100 text-xs cursor-pointer"
                    @click="()=> queryClient.invalidateQueries({ queryKey: ['containers'] })"
              >
                {{ isFetching ? "Loading..." : "Refresh" }}
              </span>
            </div>

            <div>
              <InputText class="border w-60 h-10 px-4 py-2 outline-0 mr-3"
                         placeholder="Search by container name"
                         type="text"
                         :value="containerInputText"
                         @input="handleContainerChange"
              />

              <InputText class="border w-60 h-10 px-4 py-2 outline-0 text-md"
                         placeholder="Search by image name"
                         type="text"
                         :value="imageInputText"
                         @input="handleImageChange"
              />
            </div>
          </div>

        </template>
        <Column field="Id" header="Container ID">
          <template #body="slotProps">
            <div class="cursor-pointer hover:underline"
                 @click="()=> {
                    openModal(slotProps.data)
                 }"
            >
              {{ getShortId(slotProps.data.Id) }}
            </div>
          </template>
        </Column>
        <Column field="Names" header="Name">
          <template #body="slotProps">
            <div class="cursor-pointer hover:underline"
                 @click="()=> {
                    openModal(slotProps.data)
                 }"
            >
              {{ getContainerName(slotProps.data.Names) }}
            </div>
          </template>
        </Column>
        <Column field="Image" header="Image"></Column>
        <Column field="State" header="Status">
          <template #body="slotProps">
            <Tag :value="slotProps.data.State" :severity="getSeverity(slotProps.data)"/>
          </template>
        </Column>
        <Column field="Created" header="Creation date">
          <template #body="slotProps">
            <div>
              {{ formatDate(slotProps.data.Created) }}
            </div>
          </template>
        </Column>
        <Column field="Id" header="Action">
          <template #body="slotProps">
            <div class="card flex justify-content-center">
              <Button type="button" severity="primary" icon="pi pi-ellipsis-v"
                      @click="(event) => {
                        selectedContainerId = slotProps.data.Id
                        toggle(event)
                      }"
                      aria-haspopup="true"
                      aria-controls="overlay_menu"
              />
              <Menu ref="menu" id="overlay_menu" :model="items" :popup="true" />
            </div>
          </template>
        </Column>
      </DataTable>
    </div>

    <Dialog v-if="visible" v-model:visible="visible" maximizable modal :header="getContainerName(selectedContainer.Names)"
            :style="{ width: '80rem' }"
            :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
            @hide="()=>{
              visible = false
              selectedContainerId = ''
            }"
    >
      <ContainerDetail :container="selectedContainer"/>
    </Dialog>
  </div>

</template>

<style>

</style>