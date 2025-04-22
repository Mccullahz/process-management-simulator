export namespace cmd {
	
	export class ProcessStateSnapshot {
	    time: number;
	    pid: number;
	    new: number[];
	    ready: number[];
	    running: number[];
	    waiting: number[];
	    terminated: number[];
	
	    static createFrom(source: any = {}) {
	        return new ProcessStateSnapshot(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.time = source["time"];
	        this.pid = source["pid"];
	        this.new = source["new"];
	        this.ready = source["ready"];
	        this.running = source["running"];
	        this.waiting = source["waiting"];
	        this.terminated = source["terminated"];
	    }
	}

}

