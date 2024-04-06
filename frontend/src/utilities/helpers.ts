

export const updateChartDataSet = (chart: any, metric: number, label: string) => {
    const now = new Date().toLocaleTimeString("it-IT", {hour: "2-digit", minute: "2-digit", second: "2-digit"})

    if (chart.data.datasets[0].data.length > 30) {
      chart.data.datasets[0].data.shift();
      chart.data.labels.shift();
    }
    chart.data.labels.push(now);
    chart.data.datasets[0].label = label;
    chart.data.datasets[0].data.push(metric);
    chart.update();
}
export const updateNetworkChartDataSet = (chart: any, networkInput: number, networkOutput: number, inputLabel: string, outputLabel: string) => {
    const now = new Date().toLocaleTimeString("it-IT", {hour: "2-digit", minute: "2-digit", second: "2-digit"})

    if (chart.data.datasets[0].data.length > 30 && chart.data.datasets[1].data.length > 30) {
      chart.data.datasets[0].data.shift();
      chart.data.datasets[1].data.shift();
      chart.data.labels.shift();
    }
    chart.data.labels.push(now);
    chart.data.datasets[0].label = inputLabel;
    chart.data.datasets[1].label = outputLabel;
    chart.data.datasets[0].data.push(networkInput);
    chart.data.datasets[1].data.push(networkOutput);
    chart.update();
}

export function setCpuChartData() {
    const documentStyle = getComputedStyle(document.documentElement);

    return {
        labels: [],
        datasets: [
            {
                label: 'CPU Usage (%)',
                data: [],
                fill: false,
                borderColor: documentStyle.getPropertyValue('--cyan-500'),
                tension: 0.4
            }
        ]
    };
}

export function setMemoryChartData() {
    const documentStyle = getComputedStyle(document.documentElement);

    return {
        labels: [],
        datasets: [
            {
                label: 'Memory Usage (%)',
                data: [],
                fill: false,
                borderColor: documentStyle.getPropertyValue('--orange-400'),
                tension: 0.4
            }
        ]
    };
}

export function setNetworkChartData() {
    const documentStyle = getComputedStyle(document.documentElement);

    return {
        labels: [],
        datasets: [
            {
                label: 'Network Input :',
                data: [],
                fill: false,
                borderColor: documentStyle.getPropertyValue('--cyan-500'),
                tension: 0.4
            },
            {
                label: 'Network Output :',
                data: [],
                fill: false,
                borderColor: documentStyle.getPropertyValue('--gray-500'),
                tension: 0.4
            }
        ]
    };
}

export const setChartOptions = () => {
    const documentStyle = getComputedStyle(document.documentElement);
    const textColor = documentStyle.getPropertyValue('--text-color');
    const textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
    const surfaceBorder = documentStyle.getPropertyValue('--surface-border');

    return {
        maintainAspectRatio: false,
        aspectRatio: 0.6,
        plugins: {
            legend: {
                labels: {
                    color: textColor
                }
            }
        },
        scales: {
            x: {
                options: {
                    scales: {
                        x: {
                            type: 'time',
                            time: {
                                unit: 'minute',
                                stepSize: 1 // D
                            }
                        }
                    }
                },
                ticks: {
                    color: textColorSecondary
                },
                grid: {
                    color: surfaceBorder
                }
            },
            y: {
                ticks: {
                    color: textColorSecondary,
                    min: 0,
                },
                grid: {
                    color: surfaceBorder
                }
            }
        }
    };
}
