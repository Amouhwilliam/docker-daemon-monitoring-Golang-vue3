export interface ChartDataInterface {
    labels: Array<string>;
    datasets: Array<DatasetInterface>;
}

export interface DatasetInterface {
    label: string;
    data: Array<number>;
    fill: boolean;
    borderColor: string;
    tension: number;
}