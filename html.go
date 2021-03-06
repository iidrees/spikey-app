package main

const (
	html = `<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Puppies</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.0/umd/popper.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/js/bootstrap.min.js"></script>

        <style type="text/css">
            /*
    STYLE
		*/
		@import "https://fonts.googleapis.com/css?family=Sunflower:300";

		body {
		    font-family: Sunflower, sans-serif;
		    background: #fafafa;
		}

		p {
		    font-family: Sunflower, sans-serif;
		    font-size: 1.1em;
		    font-weight: 300;
		    line-height: 1.7em;
		    color: #999;
		}

		a, a:hover, a:focus {
		    color: inherit;
		    text-decoration: none;
		    transition: all 0.3s;
		}

		#heading{
		    display: inline-block;
		    margin: 10px;
		    color: white;
		}

		.navbar {
		    padding: 15px 10px;
		    background: #56677C;
		    border: none;
		    border-radius: 0;
		    margin-bottom: 40px;
		    box-shadow: 1px 1px 3px rgba(0, 0, 0, 0.1);
		}

		.navbar-btn {
		    box-shadow: none;
		    outline: none !important;
		    border: none;
		}

		.line {
		    width: 100%;
		    height: 1px;
		    border-bottom: 1px dashed #ddd;
		    margin: 40px 0;
		}

		.nav-ul {
		    display: inline-block;
		    float: right;
		}

		.nav-li {
		    display: inline-block;
		    margin: 0px 5px;
		}



		/*
		    SIDEBAR STYLE
		*/
		.wrapper {
		    display: flex;
		}

		#sidebar {
		    min-width: 250px;
		    max-width: 250px;
		    background: #56677C;
		    color: #fff;
		    transition: all 0.3s;
		}

		#sidebar.active {
		    margin-left: -250px;
		}

		#sidebar .sidebar-header {
		    padding: 20px;
		    background: #56677C;
		}

		#sidebar ul.components {
		    padding: 20px 0;
		    border-bottom: 1px solid #47748b;
		}

		#sidebar ul p {
		    color: #fff;
		    padding: 10px;
		}

		#sidebar ul li a {
		    padding: 10px;
		    font-size: 1.1em;
		    display: block;
		}
		#sidebar ul li a:hover {
		    color: #7386D5;
		    background: #fff;
		}

		ul.CTAs {
		    padding: 20px;
		}

		ul.CTAs a {
		    text-align: center;
		    font-size: 0.9em !important;
		    display: block;
		    border-radius: 5px;
		    margin-bottom: 5px;
		}

		/*
		    CONTENT STYLE
		*/
		#content {
		    width: 100%;
		    padding: 20px;
		    min-height: 100vh;
		    transition: all 0.3s;
		}

		.rounded-circle {
		    border-radius: 50%;
		}

		.dog-image {
		    display: inline-block;
		    overflow: hidden;
		    background-position: center;
		}
		/*
		    MEDIAQUERIES
		*/
		@media (max-width: 768px) {
		    #sidebar {
		        margin-left: -250px;
		    }
		    #sidebar.active {
		        margin-left: 0;
		    }
		}

        </style>



</head>

<body>


        <div class="wrapper">
            <!-- Sidebar Holder -->
            <nav id="sidebar">
                <div class="sidebar-header">
                    <h3>SpikeySales</h3>
                </div>

                <ul class="list-unstyled components">
                    <li>
                        <a href="#">Home</a>
                    </li>
                    <li>
                        <a href="corousel.html">Images</a>
                    </li>
                    <li>
                        <a href="#">Contacts</a>
                    </li>
                </ul>

                <ul class="list-unstyled CTAs">
                    <li><a href="login.html">Login</a></li>
                    <li><a href="#">Sign Up</a></li>
                </ul>
            </nav>

            <div id="content">

                <nav class="navbar navbar-default">
                    <div class="container-fluid">

                        <div class="navbar-header">
                            <button type="button" id="sidebarCollapse" class="btn btn-info navbar-btn">
                                <i class="glyphicon glyphicon-align-left"></i>
                                <span> Menu </span>
                            </button>
                            <h2 id="heading">Spikey Sales</h2>
                        </div>

                        <div class="navbar-collapse" id="bs-example-navbar-collapse-1" style="max-width: 300px;">
                            <ul class="nav navbar-nav navbar-right nav-ul">
                                <li class="nav-li"><a href="#" style="color:white;">Home</a></li>
                                <li class="nav-li"><a href="corousel.html" style="color:white;">Images</a></li>
                                <li class="nav-li"><a href="#" style="color:white;">Contacts</a></li>
                            </ul>
                        </div>
                    </div>
                </nav>

				<button type="button" id="sidebarCollapse" class="btn-sm btn-primary ">
				<a class="btn-xs white" data-toggle="collapse" data-target=".nav-collapse" style="left: 500px">
							Spikey website details</a>
				 </button>



				<div class="nav-collapse collapse pull-left">
					<div class="row ">
						<div class="col s2">&nbsp;</div>
						<div class="col s8">
						<div class="card grey">
						<div class="card-content white-text">
						<div class="card-title">Backend that serviced this request</div>
						</div>
						<div class="card-content white">
						<table class="bordered">
						  <tbody>
							<tr>
							  <td>Name</td>
							  <td>{{.Name}}</td>
							</tr>
							<tr>
							  <td>Version</td>
							  <td>{{.Version}}</td>
							</tr>
							<tr>
							  <td>ID</td>
							  <td>{{.Id}}</td>
							</tr>
							<tr>
							  <td>Hostname</td>
							  <td>{{.Hostname}}</td>
							</tr>
							<tr>
							  <td>Zone</td>
							  <td>{{.Zone}}</td>
							</tr>
							<tr>
							  <td>Project</td>
							  <td>{{.Project}}</td>
							</tr>
							<tr>
							  <td>Internal IP</td>
							  <td>{{.InternalIP}}</td>
							</tr>
							<tr>
							  <td>External IP</td>
							  <td>{{.ExternalIP}}</td>
							</tr>
						  </tbody>
						</table>
						</div>

						</div>
						<div class="card blue">
						<div class="card-content white-text">
						<div class="card-title">Proxy that handled this request</div>
						</div>
						<div class="card-content white">
						<table class="bordered">
						  <tbody>
							<tr>
							  <td>Address</td>
							  <td>{{.ClientIP}}</td>
							</tr>
							<tr>
							  <td>Request</td>
							  <td>{{.LBRequest}}</td>
							</tr>
							<tr>
						  <td>Error</td>
						  <td>{{.Error}}</td>
							</tr>
						  </tbody>
						</table>
						</div>
						</div>
						</div>
	            <div class="col s2">&nbsp;</div>
				</div>
				</div>

	            
				<p></p>
				<p></p>	
                <p>A dog is the only thing on earth that loves you more than you love yourself.</p> 
                <p>Whoever said you can’t buy happiness forgot about puppies.</p>
                <p>Even within breeds there's enormous variety in the way a dog acts and reacts to the world around him.
                    Those differences can be due to how much he was handled as a youngster,
                    how well he was trained after bringing him home,
                    and of course the genetic luck of the draw.</p>

                <div class="line"></div>

                <div class="dog-image">
                    <img src="https://images.pexels.com/photos/1108099/pexels-photo-1108099.jpeg?cs=srgb&dl=adorable-animal-breed-1108099.jpg&fm=jpg" height="200" width="300">
                    <p> </p>
                    <h2>Golden Retiever</h2>
                <p>He is a large-sized breed of dog bred as gun dogs to retrieve shot waterfowl 
                such as ducks and upland game birds during hunting and shooting parties, 
                and were named 'retriever' because of their ability to retrieve shot game undamaged.</p>
                </div>
                    
                    
                <div class="line"></div>

                <div class="dog-image">
                <img src="https://images.pexels.com/photos/257540/pexels-photo-257540.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260" height="200" width="300">
                    <p> </p>
                </div>

                <h2>Labrador</h2>
                <p>He is a type of retriever-gun dog. A favourite disability assistance breed in many countries, 
                Labradors are frequently trained to aid the blind, those who have autism, to act as a therapy dog. 
                Additionally, they are prized as sporting and hunting dogs.</p>


                <div class="line"></div>

                <div class="dog-image">
                <img src="https://images.pexels.com/photos/208834/pexels-photo-208834.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260" height="200" width="300">
                    <p> </p>
                </div>

                <h2>German Shepherd</h2>
                <p>He is a breed of medium to large-sized working dog that originated in Germany.
                    The breed was once known as the Alsatian in Britain and Ireland.The breed is a relatively new breed of dog, 
                    with their origin dating to 1899. As part of the Herding Group, 
                    German Shepherds are working dogs developed originally for herding sheep. 
                    Since that time however, because of their strength, intelligence, trainability, and obedience, 
                    German Shepherds around the world are often the preferred breed.German Shepherd Dogs can stand as high as 26 inches at the shoulder and, 
                    when viewed in outline, presents a picture of smooth, graceful curves rather than angles.</p>

                <div class="line"></div>

                <div class="dog-image">
                <img src="https://images.pexels.com/photos/169524/pexels-photo-169524.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260" height="200" width="300">
                    <p> </p>
                </div>

                <h2>Dobermann</h2>
                <p>He is a medium-large breed of domestic dog. The Dobermann has a long muzzle and stands 
                on its toes and is not usually heavy-footed. Ideally, they have an even and graceful gait. 
                Dobermanns have markings on the chest, paws/legs, muzzle, above the eyes, and underneath the tail.</p>

                
            </div>
        </div>

        <!-- jQuery CDN -->
         <script src="https://code.jquery.com/jquery-1.12.0.min.js"></script> -
          <script type="text/javascript">
             $(document).ready(function () {
                 $('#sidebarCollapse').on('click', function () {
                     $('#sidebar').toggleClass('active');
                 });
             });
         </script>
    </body>
</html>`
)
