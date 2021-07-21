Pod::Spec.new do |spec|
  spec.name         = 'Getd'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/crypyto-panel/go-etherdata'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Etherdata Client'
  spec.source       = { :git => 'https://github.com/crypyto-panel/go-etherdata.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Getd.framework'

	spec.prepare_command = <<-CMD
    curl https://getdstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Getd.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
