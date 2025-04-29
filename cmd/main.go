package main

import (
	"log"
	"gonum.org/v1/plot"
    	"gonum.org/v1/plot/plotter"
    	"gonum.org/v1/plot/vg"
	"math"
	 "gonum.org/v1/plot/plotutil"
    	
	"github.com/healthcare-system/internal/apf"
	"github.com/healthcare-system/internal/cpt"
	"github.com/healthcare-system/pkg/blockchain"
)
func main() {
	// Initialize components
	bloomFilter := apf.NewBloomFilter(1000)
	patriciaTrie := cpt.NewPatriciaTrie()
	blockchainNetwork := blockchain.Blockchain{}

	// Real-time example: Adding five patient records to the system
	patientRecords := map[string]string{
	"patient101": "John Doe, Age: 45, Blood Type: O+, Medical History: Hypertension, Diabetes",
	"patient102": "Jane Smith, Age: 30, Blood Type: A-, Medical History: Asthma",
	"patient103": "Sam Wilson, Age: 50, Blood Type: B+, Medical History: High Cholesterol",
	"patient104": "Anna Taylor, Age: 40, Blood Type: AB+, Medical History: No Known Conditions",
	"patient105": "Mark Lee, Age: 60, Blood Type: O-, Medical History: Heart Disease",
	"patient106": "Emily Davis, Age: 55, Blood Type: A+, Medical History: Hypertension",
	"patient107": "Michael Brown, Age: 42, Blood Type: B-, Medical History: Asthma, Diabetes",
	"patient108": "Sophia Johnson, Age: 28, Blood Type: O+, Medical History: Migraine",
	"patient109": "Liam Wilson, Age: 37, Blood Type: O-, Medical History: No Known Conditions",
	"patient110": "Isabella Martinez, Age: 49, Blood Type: A+, Medical History: Thyroid Disorder",
	"patient111": "Ethan Thomas, Age: 33, Blood Type: B+, Medical History: Anxiety",
	"patient112": "Olivia Clark, Age: 25, Blood Type: AB-, Medical History: No Known Conditions",
	"patient113": "William Harris, Age: 52, Blood Type: O+, Medical History: Hypertension",
	"patient114": "Mia Lewis, Age: 34, Blood Type: A-, Medical History: Asthma",
	"patient115": "James Walker, Age: 47, Blood Type: B+, Medical History: Diabetes",
	"patient116": "Amelia Scott, Age: 50, Blood Type: O-, Medical History: Arthritis",
	"patient117": "Alexander King, Age: 39, Blood Type: AB+, Medical History: Migraine",
	"patient118": "Ella Green, Age: 29, Blood Type: A+, Medical History: No Known Conditions",
	"patient119": "Daniel Perez, Age: 56, Blood Type: B+, Medical History: High Cholesterol",
	"patient120": "Grace Turner, Age: 41, Blood Type: O+, Medical History: Hypertension",
	"patient121": "Lucas Baker, Age: 36, Blood Type: A-, Medical History: Asthma",
	"patient122": "Chloe Adams, Age: 22, Blood Type: O-, Medical History: No Known Conditions",
	"patient123": "Benjamin White, Age: 48, Blood Type: B-, Medical History: Diabetes",
	"patient124": "Zoe Hill, Age: 32, Blood Type: AB-, Medical History: No Known Conditions",
	"patient125": "Henry Moore, Age: 57, Blood Type: O+, Medical History: Hypertension",
	"patient126": "Lily Evans, Age: 27, Blood Type: A+, Medical History: No Known Conditions",
	"patient127": "Jack Taylor, Age: 43, Blood Type: B+, Medical History: High Cholesterol",
	"patient128": "Victoria Lee, Age: 38, Blood Type: O-, Medical History: Asthma",
	"patient129": "Samuel Cooper, Age: 51, Blood Type: AB+, Medical History: Diabetes",
	"patient130": "Scarlett Turner, Age: 35, Blood Type: A+, Medical History: Thyroid Disorder",
	"patient131": "Sebastian Ward, Age: 46, Blood Type: B-, Medical History: Heart Disease",
	"patient132": "Charlotte Young, Age: 31, Blood Type: O+, Medical History: Hypertension",
	"patient133": "Leo Clark, Age: 44, Blood Type: A-, Medical History: No Known Conditions",
	"patient134": "Sophia Baker, Age: 53, Blood Type: B+, Medical History: Anxiety",
	"patient135": "David Martin, Age: 40, Blood Type: O-, Medical History: No Known Conditions",
	"patient136": "Ava Parker, Age: 24, Blood Type: A+, Medical History: No Known Conditions",
	"patient137": "Mason Rivera, Age: 29, Blood Type: AB+, Medical History: High Cholesterol",
	"patient138": "Sophie Bennett, Age: 58, Blood Type: O+, Medical History: Arthritis",
	"patient139": "Jacob Kelly, Age: 34, Blood Type: A-, Medical History: Asthma",
	"patient140": "Isabella Ramirez, Age: 42, Blood Type: O+, Medical History: Thyroid Disorder",
	"patient141": "Evelyn Cox, Age: 30, Blood Type: AB+, Medical History: Anxiety",
	"patient142": "Oliver Gray, Age: 37, Blood Type: O+, Medical History: Hypertension",
	"patient143": "Emily Morris, Age: 27, Blood Type: B-, Medical History: Asthma",
	"patient144": "James Rogers, Age: 45, Blood Type: AB+, Medical History: Heart Disease",
	"patient145": "Elizabeth Reed, Age: 32, Blood Type: O+, Medical History: Migraine",
	"patient146": "Elijah Bell, Age: 41, Blood Type: A+, Medical History: No Known Conditions",
	"patient147": "Lily Powell, Age: 39, Blood Type: AB-, Medical History: High Cholesterol",
	"patient148": "Mia Brooks, Age: 46, Blood Type: B+, Medical History: Hypertension",
	"patient149": "Noah Rivera, Age: 28, Blood Type: O-, Medical History: Diabetes",
	"patient150": "Avery Richardson, Age: 54, Blood Type: AB+, Medical History: Thyroid Disorder",
	// Add more patients up to 100, following the same format
	//...
	"patient200": "Lucas Morgan, Age: 35, Blood Type: A-, Medical History: Heart Disease",
}

	// Add each patient record to Bloom Filter and Patricia Trie
	for patientID, patientData := range patientRecords {
		bloomFilter.Add(patientID)
		patriciaTrie.Insert(patientID, patientData)
		log.Printf("Added record for %s", patientID)
	}

	// Check and retrieve each patient record
	for patientID := range patientRecords {
		if bloomFilter.Check(patientID) {
			log.Printf("Patient record likely exists in the system: %s", patientID)
			if data, found := patriciaTrie.Search(patientID); found {
				log.Printf("Retrieved patient data: %s", data)
			} else {
				log.Printf("Patient data for %s not found in Patricia Trie", patientID)
			}
		} else {
			log.Printf("Patient record for %s does not exist in the system", patientID)
		}
	}

	// Initialize Blockchain Network
	blockchainNetwork.InitializeNetwork()

	if blockchainNetwork.IsInitialized() {
		log.Println("Blockchain network initialized")
	}
	/////////////////////
	  p := plot.New()

    // Set the title and axis labels
    p.Title.Text = "Scalability"

    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"

    // Create a line plotter for the sine function
    n := 100
    pts := make(plotter.XYs, n)
    for i := range pts {
        x := float64(i) / 10
        pts[i].X = x
        pts[i].Y = x*2
    }

    // Add the line to the plot
    line, err := plotter.NewLine(pts)
    if err != nil {
        panic(err)
    }
    p.Add(line)

    // Save the plot to a PNG file
    if err := p.Save(4*vg.Inch, 4*vg.Inch, "scalability.png"); err != nil {
        panic(err)
    }
	////////////////////////////

/////////////////////
	  p2 := plot.New()

    // Set the title and axis labels
    p2.Title.Text = "Security"

    p2.X.Label.Text = "X"
    p2.Y.Label.Text = "Y"

    // Create a line plotter for the sine function
    n2 := 100
    pts2 := make(plotter.XYs, n2)
    for i := range pts2 {
        x := float64(i) / 10
        pts2[i].X = x
        pts2[i].Y = x+2
    }

    // Add the line to the plot
    line, err2 := plotter.NewLine(pts2)
    if err2 != nil {
        panic(err2)
    }
    p2.Add(line)

    // Save the plot to a PNG file
    if err2 := p2.Save(4*vg.Inch, 4*vg.Inch, "security.png"); err != nil {
        panic(err2)
    }
	/////////////////////

	cosPts := make(plotter.XYs, n)
for i := range cosPts {
    x := float64(i) / 10
    cosPts[i].X = x
    cosPts[i].Y = math.Cos(x)
}

cosLine, _ := plotter.NewLine(cosPts)
cosLine.Color = plotutil.Color(1) // Change color

p.Add(line, cosLine)
	log.Println("All operations completed.")
}