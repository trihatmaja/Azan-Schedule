const vm = new Vue({
  el: '#app',
  data: {
    results: [],
    cities: [],
    addCity: false,
    city: "",
    lat: "",
    long: "",
    tz:"",
    selectedCity:"jakarta",
    cityOptions: [],
    today: [],
    currentDate: moment().format("YYYYMMDD"),
    currentMonth: moment().format("MM"),
    btnText: "Tambahkan Kota",
  },
  mounted() {
    axios.get("http://localhost:1234/api/cities/"+this.selectedCity+"/month/"+this.currentMonth)
    .then(response => {
    	this.results = response.data.data.schedule
    });

    axios.get("http://localhost:1234/api/cities")
    .then(response=> {
    	this.cities = response.data.data
    	var list=[]
		this.cities.map(function(obj) {
			list.push({text: obj.city, value: obj.city})
		});
		this.cityOptions = list
    });

    axios.get("http://localhost:1234/api/cities/"+this.selectedCity+"/date/"+this.currentDate)
    .then(response => {
    	this.today = response.data.data.schedule
    });

  },
  methods: {
    show: function (event) {
    	if(this.addCity) {
    		this.addCity = false
    		this.btnText = "Tambahkan Kota"
    	} else {
    		this.addCity = true
    		this.btnText = "Tutup"
    	}

    },
    save: function(e) {
    	data = {
    		"city": this.city,
    		"lat": parseFloat(this.lat),
    		"long": parseFloat(this.long),
    		"tz": parseFloat(this.tz),
    	}

    	axios.post('http://localhost:1234/api/generate', data)
		  .then(function (response) {
		    console.log(response);
		  })
		  .catch(function (error) {
		    console.log(error);
		  });
    },
    citychange: function(e) {
	    axios.get("http://localhost:1234/api/cities/"+this.selectedCity+"/month/"+this.currentMonth)
	    .then(response => {
	    	this.results = response.data.data.schedule
	    });
    }
  },
  watch: {
        selectedCity: function(val, oldVal){
		    axios.get("http://localhost:1234/api/cities/"+val+"/month/"+this.currentMonth)
		    .then(response => {
		    	this.results = response.data.data.schedule
    		});

    		axios.get("http://localhost:1234/api/cities/"+this.selectedCity+"/date/"+this.currentDate)
		    .then(response => {
		    	this.today = response.data.data.schedule
		    });
		        }
    }
});

Vue.filter('formatDate', function(value) {
  if (value) {
    return moment(String(value)).format('D MMMM YYYY')
  }
});