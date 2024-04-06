<script setup lang="ts">
import {onMounted, onUnmounted, ref} from "vue";
import Chart from 'primevue/chart';
import {
  setChartOptions,
  setCpuChartData,
  setMemoryChartData,
  setNetworkChartData,
  updateChartDataSet, updateNetworkChartDataSet
} from "../utilities/helpers.ts";

const {container} = defineProps(['container'])

interface ChartDataInterface {
  labels: Array<string>;
  datasets: Array<DatasetInterface>;
}

interface DatasetInterface {
  label: string;
  data: Array<number>;
  fill: boolean;
  borderColor: string;
  tension: number;
}

const eventSource = new EventSource(`/server/containers/${container.Id}/stats`);

onMounted(() => {

  eventSource.addEventListener("stats", (event) => {

    let data = JSON.parse(event.data);
    console.log(data)
    let cpuChartInstance = cpuChartRef?.value?.chart;
    let memoryChartInstance = memoryChartRef?.value?.chart;
    let networkChartInstance = networkChartRef?.value?.chart;

    let label = `Memory Usage (%) - ${data.usedMemoryMb.toFixed(2)}Mb / ${data.totalMemoryGb.toFixed(2)}Gb`
    let networkInputLabel = `Network Input: ${data.networkInputMb.toFixed(2)}Mb`
    let networkOutputLabel = `Network Output: ${data.networkOutputMb.toFixed(2)}Mb`

    updateChartDataSet(cpuChartInstance, data.cpuUsage, "CPU Usage (%)")
    updateChartDataSet(memoryChartInstance, data.memoryUsage, label)
    updateNetworkChartDataSet(networkChartInstance, data.networkInputMb, data.networkOutputMb, networkInputLabel, networkOutputLabel)
  });
  chartOptions.value = setChartOptions();
});

const cpuChartData: ChartDataInterface = setCpuChartData();
const memoryChartData: ChartDataInterface = setMemoryChartData();
const networkChartData: ChartDataInterface = setNetworkChartData();
const chartOptions = ref();
const cpuChartRef = ref<any>(null)
const memoryChartRef = ref<any>(null)
const networkChartRef = ref<any>(null)

onUnmounted(() => eventSource.close())

</script>

<template>
  <div>
    <div class="flex w-full justify-evenly">
      <Chart ref="cpuChartRef" type="line" :data="cpuChartData" :options="chartOptions" class="h-[30rem] w-1/2 mr-6"/>
      <Chart ref="memoryChartRef" type="line" :data="memoryChartData" :options="chartOptions" class="h-[30rem] w-1/2"/>
    </div>
    <div class="flex w-full justify-start mt-12">
      <Chart ref="networkChartRef" type="line" :data="networkChartData" :options="chartOptions" class="h-[30rem] w-1/2 mr-6"/>
    </div>

  </div>
</template>

<style scoped>

</style>
